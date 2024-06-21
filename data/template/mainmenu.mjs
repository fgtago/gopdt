const obj_mainmenu = document.getElementById('obj_mainmenu')
const btn_mainmenu_open = document.getElementById('btn_mainmenu_open')
const btn_mainmenu_close = document.getElementById('btn_mainmenu_close')

export async function Init() {
	console.log("initializing mainmenu module")


	window.addEventListener('click', (event)=>{
		clickOutsideMenu(event)
	})	

	obj_mainmenu.addEventListener('click', (event)=>{
		event.stopPropagation()
	})

	btn_mainmenu_open.addEventListener('click', (event)=>{
		btn_mainmenu_open_click()
		event.stopPropagation()
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

function clickOutsideMenu(event) {
	if (!obj_mainmenu.classList.contains('hidden')) {
		// main menu dibuka, close dulu
		btn_mainmenu_close_click()
	}
}
