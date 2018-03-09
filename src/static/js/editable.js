// Following line is useful when checking to see who is sending the data:
// elemToWorkOn.offsetParent.children[0].childNodes[0].firstChild.childNodes[0].id

function createField(oldText, type){
    var form = '<form class="inline-editor">';
    form += '<input type="' + type + '" class="editBox black-text" value="';
    form += oldText;
    form += '"/> </form>';
    form += '<br>';
    form += '<a href="#" class="btnSave white-text btn extra-spacing grey darken-3" id="saveButton">Save</a>';
    form += '<a href="#" class="btnDiscard btn grey darken-3" id="cancelButton">Cancel</a>';
    return form;
}

function replaceHTML(originalMethod, originalClass, elemToWorkOn, type){
    if (! $('.editing').length){
        var oldText = $(elemToWorkOn).html();
        var form = createField(oldText, type);
        $(elemToWorkOn).html('').html(form);
        document.getElementById('cancelButton').addEventListener('click', function(){
            elemToWorkOn.innerHTML = oldText;
            elemToWorkOn.classList.remove('editing');
            setTimeout(function(){
                elemToWorkOn.addEventListener('click', originalMethod);
            }, 500);
        });
        $(elemToWorkOn).unbind('click', originalMethod);
        elemToWorkOn.removeEventListener('click', originalMethod);
        document.getElementById('saveButton').addEventListener('click', function(){
            var newText = $(elemToWorkOn).parent().find('.editBox').val();
            saveChanges(originalMethod, originalClass, elemToWorkOn, newText);
        });
        $(document.body).on('submit', '.inline-editor', function(e){
            e.preventDefault();
            var newText = $(this).find('.editBox').val();
            saveChanges(originalMethod, originalClass, elemToWorkOn, newText);
        });
    }
}

function saveChanges(originalMethod, originalClass, elemToWorkOn, newText){
    elemToWorkOn.innerHTML = newText;
    elemToWorkOn.classList.remove('editing');
    setTimeout(function(){
        elemToWorkOn.addEventListener('click', originalMethod);
    }, 500);
}

function initializeEditableCards(){
    $('.textarea-editable').bind('click', function editTextArea(){
        replaceHTML(editTextArea, '.textarea-editable', this, 'textarea');
    });
    $('.textfield-editable').bind('click', function editText(){
        replaceHTML(editText, '.textfield-editable', this, 'text');
    });
    $('.timepicker-editable').bind('click', function editTimePicker(){
        replaceHTML(editTimePicker, '.timepicker-editable', this, 'timepicker');
    });
}

initializeEditableCards();
