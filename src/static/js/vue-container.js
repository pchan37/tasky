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
            handleClass: 'handle',
            dataMode: true
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
            }
        },

        showDatePicker: function(event){
            var datepicker = new MaterialDatetimePicker().on('submit', (val) => {
                var e = new Event('input', {bubbles: true});
                Vue.set(event.target, 'value', val.format('MM/DD/YYYY HH:MM:SS'));
                event.target.dispatchEvent(e);
            });
            datepicker.open();
        }
    }
});
