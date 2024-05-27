(() => {
	let MOVE_LS_KEY = '__devtool_position__'
	let activeClass = "is-active"
	let retryLimit = 30
	let retryCount = 0
	let loaded = false
	
	function getPopup(id) {
		return document.querySelector(`[data-devtool-popup="${id}"]`)
	}
	
	function closePopups(e) {
		const devtool = document.querySelector('.devtool')
		if (!devtool) {
			return
		}
		if (devtool.contains(e.target)) {
			return
		}
		for (const item of document.querySelectorAll('[data-devtool-popup]')) {
			item.classList.remove(activeClass)
		}
		for (const item of document.querySelectorAll('[data-devtool-popup-handler]')) {
			item.classList.remove(activeClass)
		}
	}
	
	function initPopups() {
		window.addEventListener('click', closePopups)
		
		const popupHandlers = document.querySelectorAll('[data-devtool-popup-handler]')
		for (const handler of popupHandlers) {
			handler.addEventListener('click', () => {
				const popup = getPopup(handler.getAttribute('data-devtool-popup-handler'))
				if (!popup) {
					return
				}
				for (const h of popupHandlers) {
					const p = getPopup(h.getAttribute('data-devtool-popup-handler'))
					if (!p || popup === p) {
						continue
					}
					h.classList.remove(activeClass)
					p.classList.remove(activeClass)
				}
				if (popup.classList.contains(activeClass)) {
					handler.classList.remove(activeClass)
					popup.classList.remove(activeClass)
				} else {
					handler.classList.add(activeClass)
					popup.classList.add(activeClass)
				}
			})
		}
	}
	
	function setInfoStatus(infoType, status) {
		const el = document.querySelector(`[data-devtool-${infoType}-build]`)
		if (!el) {
			return
		}
		if (el.classList.contains('is-ok')) {
			el.classList.remove('is-ok')
		}
		if (el.classList.contains('is-pending')) {
			el.classList.remove('is-pending')
		}
		if (el.classList.contains('is-fail')) {
			el.classList.remove('is-fail')
		}
		el.classList.add('is-'+status)
	}
	
	function initConnection() {
		if (retryCount >= retryLimit) {
			setInfoStatus('app', 'fail')
			setInfoStatus('assets', 'fail')
			return
		}
		let connection = new WebSocket(window.location.origin.replace("http", "ws") + "/_development")
		connection.onopen = function (e) {
			if (loaded) {
				window.location.href = window.location.href
			}
			loaded = true
			setInfoStatus('app', 'ok')
		}
		connection.onclose = (e) => {
			connection = null
			setInfoStatus('app', 'pending')
			setInfoStatus('assets', 'pending')
			setTimeout(initConnection, 500)
			retryCount += 1
		}
		return connection
	}
	
	function initResize() {
		const devtool = document.querySelector('.devtool')
		if (!devtool) {
			return
		}
		let debounce = null
		window.addEventListener('resize', function(e) {
			clearTimeout(debounce)
			debounce = setTimeout(() => {
				window.localStorage.removeItem(MOVE_LS_KEY)
				devtool.removeAttribute('style')
			}, 300)
		})
	}
	
	function initMovement() {
		const invertClass = 'invert'
		const devtool = document.querySelector('.devtool')
		if (!devtool) {
			return
		}
		devtool.setAttribute('style', window.localStorage.getItem(MOVE_LS_KEY))
		if ((devtool.offsetTop < (window.innerHeight) / 2)) {
			devtool.classList.add(invertClass)
		}
		
		devtool.addEventListener('mousedown', function(e) {
			e.preventDefault()
			let initialX = e.clientX
			let initialY = e.clientY
			
			function moveElement(event) {
				let currentX = event.clientX;
				let currentY = event.clientY;
				
				let deltaX = currentX - initialX;
				let deltaY = currentY - initialY;
				
				devtool.style.left = devtool.offsetLeft + deltaX + 'px';
				devtool.style.top = devtool.offsetTop + deltaY + 'px';
				
				initialX = currentX;
				initialY = currentY;
			}
			
			function stopElement(e) {
				document.removeEventListener('mousemove', moveElement);
				document.removeEventListener('mouseup', stopElement);
				window.localStorage.setItem(MOVE_LS_KEY, `top:${devtool.style.top};left:${devtool.style.left};`)
				for (const item of devtool.querySelectorAll('.devtool-button')) {
					if (item.contains(e.target)) {
						return
					}
				}
				if ((devtool.offsetTop < (window.innerHeight) / 2) && !devtool.classList.contains(invertClass)) {
					devtool.classList.add(invertClass)
				}
				if (devtool.offsetTop >= (window.innerHeight) / 2) {
					devtool.classList.remove(invertClass)
				}
			}
			
			document.addEventListener('mousemove', moveElement);
			document.addEventListener('mouseup', stopElement);
		})
	}
	
	if (window.devtool) {
		initPopups()
		initMovement()
		initResize()
	}
	
	document.addEventListener('DOMContentLoaded', () => {
		initPopups()
		initConnection()
		initMovement()
		initResize()
		window.devtool = true
	})
})()