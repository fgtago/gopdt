import * as webmask from './mask.mjs'


export async function Prepare() {
	console.log('Preparing module')
	window.$mask = webmask

	// create_mask()
}

export async function Init() {
	console.log('Default init')
}

export async function Ready() {
	console.log('Module Ready')

	var masks = document.getElementsByClassName('pagecover-mask')
	for (var elm of masks) {
		elm.parentNode.removeChild(elm)
	}
}

