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
        tasks: [
            {title: 'Dishes', body: 'Do the dishes', time: '03/20/2018 19:03:00', editing: true},
            {title: 'Trash', body: 'Take out the trash', time: '03/20/2018 19:03:00', editing: false},
            {title: 'Lawn', body: 'Mow the lawn', time: '03/20/2018 19:03:00', editing: false}
        ]
    },
    mounted: function() {
        lmdd.set(document.getElementById('drag-scope'), {
            containerClass: 'todo-item',
            draggableItemClass: 'card',
            dataMode: true,
            dragstartTimeout: 175
        });

        this.$el.addEventListener('lmddend', this.handleDragEvent);
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
            this.tasks.unshift({title: '', body: '', time: '', editing: true});
            console.log(this.tasks);
        },

        showDatePicker: function(event){
            var datepicker = new MaterialDatetimePicker().on('submit', (val) => {
                var e = new Event('input', {bubbles: true});
                Vue.set(event.target, 'value', val.format('MM/DD/YYYY HH:MM:SS'));
                event.target.dispatchEvent(e);
            });
            datepicker.open();
        },

        newTask: function(event, task, index){
            this.sendTaskRequest(event, '/new_task',
                                 {'index': index, 'title': task.title, 'time': task.time, 'body': task.body});
        },

        updateTask: function(event, task, index){
            task.editing = false;
            this.sendTaskRequest(event, '/update_task',
                                 {'index': index, 'title': task.title, 'time': task.time, 'body': task.body});
        },

        updateTaskPosition: function(event, startIndex, endIndex){
            this.sendTaskRequest(event, '/update_task_pos',
                                 {'startIndex': startIndex, 'endIndex': endIndex});
        },

        deleteTask: function(event, tasks, index){
            tasks.splice(index, 1);
            this.sendTaskRequest(event, '/delete_task', {'startIndex': index, 'endIndex': -1});
        },

        sendTaskRequest: function(event, url, data){
            var XHR = new XMLHttpRequest();
            var urlEncodedData = JSON.stringify(data);

            XHR.open('POST', url);
            XHR.setRequestHeader('Content-Type', 'application/json');
            XHR.send(urlEncodedData);
        }
    }
});
