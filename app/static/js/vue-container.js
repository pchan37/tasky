(function() {
    var taskQueue = new Queue();

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
            handleDragEvent: async function(event) {
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
                    const success = await this.updateTaskPosition(event, startIndex, endIndex);
                    if (success) {
                        this.tasks.splice(endIndex, 0, this.tasks.splice(startIndex, 1)[0]);
                    }
                }
            },

            addNewTaskForm: function() {
                this.tasks.unshift({title: '', body: '', time: '', editing: true, action: 'new'});
            },

            showDateTimePicker: function(event) {
                event.stopPropagation();
                var initialInput = event.target.value;
                var time = initialInput.substring(11);
                M.Datepicker.init(event.target, {
                    format: 'mm/dd/yyyy',
                    onClose: function() {
                        if (datepicker.toString() !== '') {
                            event.target.value = '';
                            timepicker.open();
                        }
                    }
                });
                var datepicker = M.Datepicker.getInstance(event.target);
                datepicker.close();

                M.Timepicker.init(event.target, {
                    defaultTime: time
                });
                var timepicker = M.Timepicker.getInstance(event.target);
                timepicker.modal.options.onOpenStart = function() {
                    event.target.value = initialInput;
                };
                timepicker.modal.options.onCloseEnd = function() {
                    if (typeof timepicker.time !== "undefined") {
                        event.target.value = datepicker.toString() + ' ' + event.target.value;
                        event.target.dispatchEvent(new Event('change'));
                        event.target.blur();
                        event.target.addEventListener('focus', function(event) {
                            vue.showDateTimePicker(event);
                        }, {once: true});
                    }
                };
                timepicker.close();

                datepicker.open();
            },

            loadTasks: function() {
                fetch('/load_tasks', {credentials: "same-origin"}).then(function(response) {
                    response.text().then(function(text) {
                        var jsonData = JSON.parse(text);
                        for(var i = 0; i < jsonData.length; i++) {
                            jsonData[i]['editing'] = false;
                        }
                        vue.tasks = jsonData;
                    });
                });
            },

            newTask: async function(event, task) {
                queueTask('/new_task',
                          {'index': 0, 'title': task.title, 'time': task.time, 'body': task.body});

                const success = await this.sendTask();
                if (success) {
                    task.editing = false;
                    task.action = '';
                }
            },

            updateTask: async function(event, task, index) {
                queueTask('/update_task',
                          {'index': index, 'title': task.title, 'time': task.time, 'body': task.body});

                const success = await this.sendTask();
                if (success) {
                    task.editing = false;
                }
            },

            updateTaskPosition: async function(event, startIndex, endIndex) {
                queueTask('/update_task_pos',
                          {'startIndex': startIndex, 'endIndex': endIndex});

                const success = await this.sendTask();
                return success;
            },

            deleteTask: async function(event, tasks, index) {
                queueTask('/delete_task', {'startIndex': index, 'endIndex': -1});

                const success = await this.sendTask();
                if (success) {
                    tasks.splice(index, 1);
                }
            },

            sendTask: async function() {
                var task = taskQueue.dequeue();
                const result = await sendTaskRequest(task.url, task.data).then(() => {
                    return true;
                }).catch(() => {
                    alert('Error sending data to server, ensure that you have internet access!');
                    return false;
                });
                return result;
            }
        }
    });

    function newTask(url, data) {
        return {
            'url': url,
            'data': data
        };
    }

    function queueTask(url, data) {
        var task = newTask(url, data);
        taskQueue.enqueue(task);
    }

    function sendTaskRequest(url, data) {
        return new Promise((resolve, reject) => {
            var XHR = new XMLHttpRequest();
            var urlEncodedData = JSON.stringify(data);

            XHR.open('POST', url);
            XHR.setRequestHeader('Content-Type', 'application/json');
            XHR.send(urlEncodedData);
            XHR.onreadystatechange = function() {
                if (XHR.readyState === 4 && XHR.status === 200) {
                    resolve(XHR.status);
                } else if (XHR.readyState === 0 || XHR.readyState === 4) {
                    reject(new Error('Data not sent successfully!'));
                }
            };
        });
    }

    function notifyUserOffline() {
        alert('You have been disconnected.  Please ensure that you have internet connection before proceeding!');
    }

    window.onload = function() {
        history.replaceState('', '', '/');
    };
    window.addEventListener('offline', notifyUserOffline);

})();
