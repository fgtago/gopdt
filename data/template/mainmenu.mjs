const obj_mainmenu = document.getElementById('obj_mainmenu')
const btn_mainmenu_open = document.getElementById('btn_mainmenu_open')
const btn_mainmenu_close = document.getElementById('btn_mainmenu_close')

export async function Init() {
	console.log("initializing mainmenu module")

	btn_mainmenu_open.addEventListener('click', ()=>{
		btn_mainmenu_open_click()
	})

	btn_mainmenu_close.addEventListener('click', ()=>{
		btn_mainmenu_close_click()
	})
}

function btn_mainmenu_open_click() {
	obj_mainmenu.classList.remove('hidden')
}

function btn_mainmenu_close_click() {
	obj_mainmenu.classList.add('hidden')
}


