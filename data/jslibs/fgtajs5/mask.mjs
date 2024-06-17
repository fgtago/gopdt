let zIndexStart = 1000;

export function show(message, buttons) {

	var masks = document.getElementsByClassName('page-cover-mask')
	var zIndex =  zIndexStart + masks.length + 1
	var elName = `page-cover-mask-${zIndex}`

	let elCMask = document.createElement('div')
	elCMask.classList.add('page-cover-mask')
	elCMask.style.zIndex = zIndex
	elCMask.id = elName
	elCMask.remove = () => {
		elCMask.parentNode.removeChild(elCMask);
		elCMask = null;
	}


	if (typeof message==='string') {
		let elCMsg = document.createElement('div')
		elCMsg.classList.add('page-cover-mask-message')

		if (buttons!==undefined) {
			var divMsg = document.createElement('div')
			divMsg.innerHTML = message

			var divButtonContainer = document.createElement('div')
			divButtonContainer.classList.add('page-cover-mask-message-buttoncontainer')

			elCMask.buttons = {}
			for (var btnid in buttons) {
				let opt = buttons[btnid]
				let elbtn = document.createElement('button')
				elbtn.classList.add('page-cover-mask-message-button')
				elbtn.innerHTML = opt.text
				elbtn.addEventListener('click', (evt)=>{
					if (typeof opt.click==='function') {
						opt.click(elCMask, evt)
					} else {
						elCMask.remove()
					}
				})
				divButtonContainer.appendChild(elbtn)
			}
			
			elCMsg.appendChild(divMsg)
			elCMsg.appendChild(divButtonContainer)

		} else {
			elCMsg.innerHTML = message
		}
		elCMask.appendChild(elCMsg)
	}

	document.body.appendChild(elCMask)


	return elCMask;
}


export function info(message) {
	return show(message, {
		btn_ok: {
			text: 'Ok',
		}
	})
}

export function error(message) {
	var errormessage = `<div class="page-cover-mask-message-error-title">Error</div>
		<div>
			${message}
		</div>
	`

	return show(errormessage, {
		btn_ok: {
			text: 'Ok',
		}
	})
}



/** contoh penggunaan
 * 
	// waiting
	let mask = window.$mask.show('wait...')

	// asking
	let mask = window.$mask.show('ok apa cancel ?', {
		btn_ok: {
			text: 'Ok',
			click: (elmask, evt) => {
				console.log('button click')
				elmask.remove()
			}
		},

		btn_cancel: {
			text: 'Cancel'
		}
	})





 * 
 */