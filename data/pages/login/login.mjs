const obj_email = document.getElementById('login-obj_email')
const obj_password = document.getElementById('login-obj_password')
const obj_rememberme = document.getElementById('login-obj_rememberme')
const form = document.getElementById('login-form')


export async function Init() {
	console.log('initializing login module')


	obj_email.addEventListener('keypress', (event)=>{
		if (event.code === "Enter") {
			obj_email_enter(event)
		}
	})


	obj_password.addEventListener('keypress', (event)=>{
		if (event.code === "Enter") {
			obj_password_enter(event)
		}	
	})

	obj_email.focus()
	obj_email.select()

}


export function Submit(form) {
	return false
}


function obj_email_enter(event) {
	if (obj_email.value.trim() !== '') {
		obj_password.focus()
	}
}


function obj_password_enter(event) {
	if (obj_password.value.trim() !== '') {
		form.submit()
	}
}
