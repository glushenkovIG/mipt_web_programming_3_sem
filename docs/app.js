function GetTextFromTask(Input){
	if (!Input.value) {
		return 'Сделать задание номер #3 по web-программированию';
	};
	return Input.value;
}

function AddTask(){

	var new_task = document.createElement('li');
	new_task.id = 'task';
	ul.appendChild(new_task);

	var task_text = document.createElement('span');
		new_task.appendChild(task_text);
		task_text.innerHTML = GetTextFromTask(input);
		task_text.id = 'task_text';


	var del_new_task_button = document.createElement('button');

		del_new_task_button.id = 'del_task';
			new_task.appendChild(del_new_task_button);
			del_new_task_button.innerHTML = "Удалить";

		del_new_task_button.addEventListener("click", function(){
			ul.removeChild(new_task);
		});
}


var root = document.getElementById('root');

var ul = document.createElement('ul');
	root.appendChild(ul);
	ul.id = "my_ul";

var input = document.createElement('input')
	input.id = "add_task_input";
	input.name = 'text_of_task'
	root.appendChild(input);

var add_task_button = document.createElement('button');
	add_task_button.innerHTML = "Добавить";
	add_task_button.id = "add_task";
	root.appendChild(add_task_button);
	add_task_button.onclick = AddTask;
	AddTask();
