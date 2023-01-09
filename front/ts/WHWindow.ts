import { _sgQuery as $ } from "./Query.js";

import { getParentElement, fire } from "./Tool.js";

import { Float, Move, Drag, Resize } from "./Float.js";

class WHWinIcon extends HTMLElement {
	constructor() {
		super();

	}
	public connectedCallback() {
		this.slot = "icon";
	}
}
customElements.define("wh-win-icon", WHWinIcon);

class WHWinCaption extends HTMLElement {
	constructor() {
		super();
	}

	public connectedCallback() {
		this.slot = "caption";
	}
}
customElements.define("wh-win-caption", WHWinCaption);

class WHWinBody extends HTMLElement {
	constructor() {
		super();

	}

	public connectedCallback() {
		this.slot = "body";
	}
}
customElements.define("wh-win-body", WHWinBody);

class WHWinHeader extends HTMLElement {
	public win: WHWin = null;
	public mode = "";

	static get observedAttributes() {
		return ["mode", "paz"];
	}

	constructor() {
		super();

		const template = document.createElement("template");

		template.innerHTML = `
			
			<link rel="stylesheet" href="./../css/WHHeader.css">
			
				<div class="icon"><slot name="icon"></slot></div>
				<div class="caption"><slot name="caption"></slot></div>
				<div class="option"><slot name="option"></slot></div>
				<div class="control">
					<button class="min"></button>
					<button class="auto"></button>
					<button class="max"></button>
					<button class="exit"></button>
				</div>
			
			
			
			<slot></slot>

		`;

		this.attachShadow({ mode: "open" });

		this.shadowRoot.appendChild(template.content.cloneNode(true));

		/*this.shadowRoot.addEventListener("slotchange", (event) => {
			[...event.target.assignedElements()].forEach(e => {

			})

		});*/


	}



	public connectedCallback() {

		this.slot = "header";

		$(this.shadowRoot).q(".min").on("click", event => {
			this.getWin().setAttribute("mode", "min");
		});

		$(this.shadowRoot).q(".max").on("click", event => {
			this.getWin().setAttribute("mode", "max");
		});

		$(this.shadowRoot).q(".auto").on("click", event => {
			this.getWin().setAttribute("mode", "");
			//this.getWin().updatePos();
			this.getWin().updateSize();
		});

		$(this.shadowRoot).q(".exit").on("click", event => {
			this.getWin().setAttribute("visibility", "hidden");
		});

		$(this).on("dblclick", event => {
			if (this.getWin().mode === "max") {
				this.getWin().setAttribute("mode", "");
			} else {
				this.getWin().setAttribute("mode", "max");
			}
		})

		this.getWin().addEventListener("mode-changed", (e: CustomEvent) => {
			this.setAttribute("mode", e.detail.mode);
		})
	}

	public attributeChangedCallback(name, oldVal, newVal) {
		console.log("attributeChangedCallback");


	}

	public getWin() {
		return getParentElement(this, "wh-win") as WHWin;
	}


}

customElements.define("wh-win-header", WHWinHeader);


export class WHWin extends HTMLElement {


	static get observedAttributes() {
		return ["visibility", "mode", "resizable", "movible", "left", "top", "right", "bottom", "width", "height"];
	}

	constructor() {
		super();

		const template = document.createElement("template");

		template.innerHTML = `

			<link rel="stylesheet" href="./../css/WHWindow.css">
			
			<slot name="header"></slot>
			<slot name="body"></slot>
			<slot></slot>
			`;

		this.attachShadow({ mode: "open" });

		this.shadowRoot.appendChild(template.content.cloneNode(true));
		/*
		this.shadowRoot.addEventListener("slotchange", (event) => {
					[...event.target.assignedElements()].forEach(e => {
						console.log(e.nodeName)
						if (e.nodeName.toLowerCase() == "wh-win-header") {
							e.slot = "header"
						}

						if (e.nodeName.toLowerCase() == "wh-win-body") {
							e.slot = "body"
						}
					})

				});
		*/

	}



	public connectedCallback() {
		const holder = this.querySelector(`wh-win-header`) as HTMLElement;

		Float.init(this);
		Move.init({
			main: this, hand: holder, onDrag: (info) => {
				if (this.mode === "max") {

					const w = this.offsetWidth;
					this.setAttribute("mode", "custom");
					const w2 = this.offsetWidth;
					this.style.left = (info.x - (w2 * (info.x - info.left) / w)) + "px";

					return true;

				}

			}
		});

	}

	public disconnectedCallback() {
		console.log("disconnectedCallback");
	}

	public attributeChangedCallback(name, oldVal, newVal) {
		console.log("attributeChangedCallback");

		switch (name) {
			case "mode":
				fire(this, "mode-changed", { mode: newVal });
				this._setMode();
				break;
			case "left":
			case "right":
			case "bottom":
			case "top":
				this.updatePos();
				break;
			case "width":
			case "height":
				this.updateSize();
				break;
			case "resizable":
				this._resizable();
			case "visibility":
				


		}
		Float.upIndex(this);
	}

	set mode(value) {
		if (Boolean(value)) {
			this.setAttribute("mode", value);
		} else {
			this.removeAttribute("mode");
		}
	}

	get mode() {
		return this.getAttribute("mode");
	}


	set width(value) {
		if (Boolean(value)) {
			this.setAttribute("width", value);
		} else {
			this.removeAttribute("width");
		}
	}

	get width() {
		return this.getAttribute("width");
	}

	set height(value) {
		if (Boolean(value)) {
			this.setAttribute("height", value);
		} else {
			this.removeAttribute("height");
		}
	}

	get height() {
		return this.getAttribute("width");
	}

	set left(value) {
		if (Boolean(value)) {
			this.setAttribute("left", value);
		} else {
			this.removeAttribute("left");
		}
	}

	get left() {
		return this.getAttribute("left");
	}

	set right(value) {
		if (Boolean(value)) {
			this.setAttribute("right", value);
		} else {
			this.removeAttribute("right");
		}
	}

	get right() {
		return this.getAttribute("right");
	}

	set top(value) {
		if (Boolean(value)) {
			this.setAttribute("top", value);
		} else {
			this.removeAttribute("top");
		}
	}

	get top() {
		return this.getAttribute("top");
	}

	set bottom(value) {
		if (Boolean(value)) {
			this.setAttribute("bottom", value);
		} else {
			this.removeAttribute("bottom");
		}
	}

	get bottom() {
		return this.getAttribute("bottom");
	}

	set visibility(value) {
		if (Boolean(value)) {
			this.setAttribute("visibility", value);
		} else {
			this.removeAttribute("visibility");
		}
	}

	get visibility() {
		return this.getAttribute("visibility");
	}

	set resizable(value) {
		if (Boolean(value)) {
			this.setAttribute("resizable", "");
		} else {
			this.removeAttribute("resizable");
		}
	}

	get resizable() {
		return this.hasAttribute("resizable");
	}

	public test() {
		alert("soy tu padre");
	}
	public updatePos() {
		if(this.hasAttribute("left")){
			this.style.left = this.left;
		}
		if(this.hasAttribute("right")){
			this.style.right = this.right;
		}
		if(this.hasAttribute("top")){
			this.style.top = this.top;
		}
		if(this.hasAttribute("bottom")){
			this.style.bottom = this.bottom;
		}

		
	}
	public updateSize() {
		this.style.width = this.width;
		this.style.height = this.height;
	}

	_setMode() {
		if (this.mode === "modal") {
			const header = this.querySelector(`wh-win-header`);
			if (header) {
				header.mode = "modal";
			}

		}
	}

	_resizable() {
		if (this.resizable) {
			Resize.init({
				main: this,
				onStart: (info) => {
					this.style.left = info.left + "px";
					this.style.width = info.width + "px";
					this.style.top = info.top + "px";
					this.style.height = info.height + "px";
					this.setAttribute("mode", "custom");
				},
				onRelease: (info) => {
					//this.width = info.width + "px";
					//this.height = info.height + "px";
				}

			});
		} else {

		}
	}

}

customElements.define("wh-win", WHWin);