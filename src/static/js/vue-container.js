Vue.component('todo-container', {
    template: '<div class="todo-container"><slot></slot></div>',
    props: ['data']
});

Vue.component('todo-item', {
    props: ['data']
});

var vue = new Vue({
    el: '#vue-container',
    data: {
        tasks: [
            ['Dishes', 'Do the dishes', '2hrs left'],
            ['Trash', 'Take out the trash', '2hrs left'],
            ['Lawn', 'Mow the lawn', '2hrs left'],
        ]
    },
    mounted: function() {
        lmdd.set(document.getElementById('drag-scope'), {
            containerClass: 'todo-container',
            draggableItemClass: 'card',
            handleClass: 'handle',
            dataMode: true
        });

        this.$el.addEventListener('lmddend', this.handleDragEvent);
    },
    methods: {
        handleDragEvent: function(event) {
            console.log('hi');
            var newIndex = event.detail.to.index;
            var oldIndex = event.detail.from.index;
            var newContainer = event.detail.to.container.__vue__.data;
            var oldContainer = event.detail.from.container.__vue__.data;
            if (event.detail.dragType === 'move') {
                newContainer.splice(newIndex, 0, oldContainer.splice(oldIndex, 1)[0]);
            }
        }//,
        // addNewTodo: function() {
        //     this.todos.regular.push(this.newTodoText)
        //     this.newTodoText = ''
        // }
    }
});
