(function() {
    var taskQueue = new Queue();
    var isRunning = false;

    Vue.component('todo-container', {
        template: '<div class="todo-container"><slot></slot></div>',
        props: ['data']
    });

    Vue.component('todo-item', {
        template: '<div class="todo-item"><slot></slot></div>',
        props: ['data']
    });

    var vue = new Vue({
        el: '#vue-container',
        data: {
            tasks: []
        },
        mounted: function() {
            lmdd.set(document.getElementById('drag-scope'), {
                containerClass: 'todo-item',
                draggableItemClass: 'card',
                dataMode: true,
                dragstartTimeout: 175
            });

            this.$el.addEventListener('lmddend', this.handleDragEvent);
            this.loadTasks();
        },
        methods: {
            handleDragEvent: function(event) {
                var fromContainer = event.detail.from.container;
                var toContainer = event.detail.to.container;
                var startIndex = this.tasks.length - 1;
                while (fromContainer.nextElementSibling != null){
                    startIndex--;
                    fromContainer = fromContainer.nextElementSibling;
                }
                var endIndex = this.tasks.length - 1;
                while (toContainer.nextElementSibling != null){
                    endIndex--;
                    toContainer = toContainer.nextElementSibling;
                }
                if (event.detail.dragType === 'move') {
                    this.tasks.splice(endIndex, 0, this.tasks.splice(startIndex, 1)[0]);
                    this.updateTaskPosition(event, startIndex, endIndex);
                }
            },

            addNewTaskForm: function() {
                this.tasks.unshift({title: '', body: '', time: '', editing: true, action: 'new'});
            },

            showDatePicker: function(event) {
                var datepicker = new MaterialDatetimePicker().on('submit', (val) => {
                    var e = new Event('input', {bubbles: true});
                    Vue.set(event.target, 'value', val.format('MM/DD/YYYY HH:MM:SS'));
                    event.target.dispatchEvent(e);
                });
                datepicker.open();
            },

            loadTasks: function() {
                fetch('/load_tasks').then(function(response) {
                    response.text().then(function(text) {
                        var jsonData = JSON.parse(text);
                        for(var i = 0; i < jsonData.length; i++) {
                            jsonData[i]['editing'] = false;
                        }
                        vue.tasks = jsonData;
                    });
                });
            },

            newTask: function(event, task) {
                queueTask(event, '/new_task',
                          {'index': 0, 'title': task.title, 'time': task.time, 'body': task.body});
                this.sendTask();
                task.editing = false;
            },

            updateTask: function(event, task, index) {
                queueTask(event, '/update_task',
                          {'index': index, 'title': task.title, 'time': task.time, 'body': task.body});
                this.sendTask();
                task.editing = false;
            },

            updateTaskPosition: function(event, startIndex, endIndex) {
                queueTask(event, '/update_task_pos',
                          {'startIndex': startIndex, 'endIndex': endIndex});
                this.sendTask();
            },

            deleteTask: function(event, tasks, index) {
                queueTask(event, '/delete_task', {'startIndex': index, 'endIndex': -1});
                this.sendTask();
                tasks.splice(index, 1);
            },

            sendTask: function(overrideIsRunning=false) {
                if ((!overrideIsRunning && isRunning) || taskQueue.isEmpty()) {
                    isRunning = false;
                    return;
                }

                isRunning = true;
                var task = taskQueue.dequeue();
                let promise = new Promise(
                    (resolve, reject) => {
                        resolve(sendTaskRequest(task.event, task.url, task.data));
                    }
                );
                promise.then(
                    function() {
                        console.log('successful!');
                        vue.sendTask(true);
                    }
                ).catch((err) => {
                    alert('Error sending data to server, ensure that you have internet access!');
                });
            }

        }
    });

    function newTask(event, url, data) {
        return {
            'event': event,
            'url': url,
            'data': data
        };
    }

    function queueTask(event, url, data) {
        var task = newTask(event, url, data);
        taskQueue.enqueue(task);
    }

    function sendTaskRequest(event, url, data) {
        var XHR = new XMLHttpRequest();
        if (XHR.overrideMimeType)
            XHR.overrideMimeType("text/json");
        var urlEncodedData = JSON.stringify(data);

        XHR.open('POST', url);
        XHR.setRequestHeader('Content-Type', 'application/json');
        XHR.send(urlEncodedData);
    }

    window.onload = function() {
        history.replaceState("", "", "/");
    };

})();
