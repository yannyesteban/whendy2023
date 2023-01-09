import { loadScript } from "./LoadScript.js";
import { Q as $ } from "./Q.js";
import { getParentElement, fire } from "./Tool.js";

class MapboxMark extends HTMLElement {
    _marker = null;
    _infowindow = null;
    _popup = null;
    constructor() {
        super();
        const template = document.createElement("template");

        template.innerHTML = `
            <style>
            :host {
            }
            slot{
                display:none
            }
            </style>
            <slot></slot>

        `;

        this.attachShadow({ mode: "open" });

        this.shadowRoot.appendChild(template.content.cloneNode(true));

        const slot = this.shadowRoot.querySelector("slot");

        slot.addEventListener("slotchange", (e) => {


            const nodes = slot.assignedElements();

            //this._infowindow.setHTML(this.innerHTML);
            //infowindow.setDOMContent(this._popup);
            //console.log(nodes)
            //this._popup.innerHTML  = nodes[0].innerHTML;
            //this._popup.data = this.info;
        });

    }

    static get observedAttributes() {
        return ["latitude", "longitude", "heading", "image", "icon", "info", "visible", "follow"];
    }

    public connectedCallback() {

        this.setAttribute("role", "mark");
        let el = document.createElement("img");
        el.className = "marker";

        el.src = this.icon;
        //el.style.width = this.width;
        el.style.height = "24px";
        this._marker = new mapboxgl.Marker(el)
            .setLngLat([Number(this.longitude), Number(this.latitude)])
            .addTo(this.getMapApi());




        const markerHeight = 24, markerRadius = 0, linearOffset = 0;
        const popupOffsets = {
            "top": [0, 0],
            "top-left": [0, 0],
            "top-right": [0, 0],
            "bottom": [0, -markerHeight / 2 + 5],
            "bottom-left": [linearOffset, (markerHeight - markerRadius + linearOffset) * -1],
            "bottom-right": [-linearOffset, (markerHeight - markerRadius + linearOffset) * -1],
            "left": [markerRadius, (markerHeight - markerRadius) * -1],
            "right": [-markerRadius, (markerHeight - markerRadius) * -1]
        };

        this._infowindow = new mapboxgl.Popup({
            offset: popupOffsets,
            className: "my-class"
        }).setMaxWidth("300px");

        this._marker.setPopup(this._infowindow);
    }

    public disconnectedCallback() {
        console.log("disconnectedCallback");
        this._marker.remove();

    }

    public attributeChangedCallback(name, oldVal, newVal) {
        console.log("attributeChangedCallback");

        switch (name) {
            case "latitude":
                this._setPosition();
                break;
            case "longitude":
                this._setPosition();
                break;
            case "heading":
                this._setHeading();
                break;

            case "image":
                break;

            case "icon":
                this._setIcon();
                break;

            case "visible":
                this._visible();
                break;

            case "trace":
                break;

            case "follow":
                break;
        }


    }

    set icon(value) {
        if (Boolean(value)) {
            this.setAttribute("icon", value);
        } else {
            this.removeAttribute("icon");
        }
    }

    get icon() {
        return this.getAttribute("icon")
    }

    set name(value) {
        this.setAttribute("name", value);
    }

    get name() {
        return this.getAttribute("name");
    }

    set heading(value) {
        if (Boolean(value)) {
            this.setAttribute("heading", value);
        } else {
            this.removeAttribute("heading");
        }
    }

    get heading() {
        return this.getAttribute("heading");
    }

    set latitude(value) {
        this.setAttribute("latitude", value);
    }

    get latitude() {
        return this.getAttribute("latitude");
    }


    set longitude(value) {
        this.setAttribute("longitude", value);
    }

    get longitude() {
        return this.getAttribute("longitude");
    }


    set visible(value) {

        if (Boolean(value)) {
            this.setAttribute("visible", "");
        } else {
            this.removeAttribute("visible");
        }
    }



    get visible() {
        return this.hasAttribute("visible");
    }



    get follow() {
        return this.hasAttribute("follow");
    }


    set follow(value) {

        if (Boolean(value)) {
            this.setAttribute("follow", "");
        } else {
            this.removeAttribute("follow");
        }
    }


    get trace() {
        return this.hasAttribute("trace");
    }


    set trace(value) {

        if (Boolean(value)) {
            this.setAttribute("trace", "");
        } else {
            this.removeAttribute("trace");
        }
    }

    _visible() {
        if (!this.visible) {

        }

    }

    set mark(info) {


    }

    public getMap(): MapboxMaps {
        //const map = getParentElement(this, "google-maps");

        return getParentElement(this, "mapbox-maps");
    }

    public getMapApi() {
        const map = this.getMap();
        if (map) {
            return map.getApi();
        }


        return null;
    }

    set info(data) {

        console.log(data)
        const popup = document.createElement("wh-info");
        popup.template = this.innerHTML;
        popup.data = data;
        popup.mode = "ready";

        this._infowindow.setDOMContent(popup);
    }

    set update(info) {
        for (let x in info) {
            this[x] = info[x];
        }
    }

    _setPosition() {
        //this.longitude = longitude;
        //this.latitude = latitude;
        const latLng = { lat: Number(this.latitude), lng: Number(this.longitude) };
        this._marker.setLngLat([Number(this.longitude), Number(this.latitude)]);
        console.log("position", latLng)
        return;
        if (this.activeFollow) {
            console.log("ok")
            //this.panTo();
        }
    }

    _setIcon() {
        this._marker.getElement().src = this.icon;
    }

    _setHeading() {
        this._marker.setRotation(this.heading);
    }

    panTo() {
        const map = this.getMap();
        if (map) {
            map.panTo({ latitude: Number(this.latitude), longitude: Number(this.longitude) });
        }
    }

    flyTo(zoom: number) {
        const map = this.getMap();
        if (map) {
            map.flyTo({ latitude: Number(this.latitude), longitude: Number(this.longitude), zoom });
        }
    }
}

customElements.define("mapbox-mark", MapboxMark);


export class MapboxMaps extends HTMLElement {

    public apiKey = "AIzaSyBhPsH8OjHCypjgwt_Dl7A_W8wlBbyPink";
    public apiUrl = "https://maps.google.com/maps/api/js?key=AIzaSyCr8MljMe17YC07PuG9CtOdHSZDZgAvmew&libraries=drawing";

    //public longitude: number = -66.903603;
    //public latitude: number = 10.480594;

    #config = {};
    #map = null;

    static scriptLoaded = false;

    static loadApiFile() {

        return new Promise((resolve, reject) => {

            const key = "AIzaSyBhPsH8OjHCypjgwt_Dl7A_W8wlBbyPink";
            const url = "https://api.mapbox.com/mapbox-gl-js/v2.7.0/mapbox-gl.js";

            loadScript(url, true)
                .then(message => {
                    console.log(message);

                    MapboxMaps.scriptLoaded = true;
                    mapboxgl.accessToken = "pk.eyJ1IjoieWFubnkyNCIsImEiOiJjazYxZnM5dzMwMzk1M21xbjUyOHVmdjV0In0.4ifqDgs5_PqZd58N1DcVaQ";
                    resolve({
                        status: true
                    });

                })
                .catch(message => {
                    MapboxMaps.scriptLoaded = false;
                    reject({
                        status: true
                    });

                })


        });




    }

    static get observedAttributes() {
        return ["type"];
    }

    constructor() {
        super();
        return;
        const template = document.createElement("template");

        template.innerHTML = `
            <style>
            :host {
            }
            </style>
            <slot></slot>

        `;

        this.attachShadow({ mode: "open" });

        this.shadowRoot.appendChild(template.content.cloneNode(true));

        const slot = this.shadowRoot.querySelector("slot");

        slot.addEventListener("slotchange", (e) => {
            console.log(9999)
            //const nodes = slot.assignedNodes();
        });
    }

    
    connectedCallback() {
        console.log("connectedCallback");

        this.init();


    }

    public disconnectedCallback() {
        console.log("disconnectedCallback");
        

    }


    attributeChangedCallback(name, oldVal, newVal) {
        console.log("attributeChangedCallback");
        switch (name) {
            case "longitude":
                //this.render();
                break;
            case "latitude":
                //this.render();
                break;
        }

    }

    set latitude(value) {
        this.setAttribute("latitude", value);
    }

    get latitude() {
        return this.getAttribute("latitude");
    }


    set longitude(value) {
        this.setAttribute("longitude", value);
    }

    get longitude() {
        return this.getAttribute("longitude");
    }

    set zoom(value) {
        this.setAttribute("zoom", value);
    }

    get zoom() {
        return this.getAttribute("zoom");
    }


    public init(message?) {

        if (!MapboxMaps.scriptLoaded) {
            MapboxMaps.loadApiFile()
                .then(() => {
                    this.render();
                }).catch(() => {
                    alert("error");
                });
            return;
        }

        this.render();


    }

    public render() {

        const map = new mapboxgl.Map({
            style: 'mapbox://styles/mapbox/streets-v11', // style URL
            container: this,
            zoom: this.zoom || 10,
            center: [this.longitude, this.latitude],
            attributionControl: false
        });

        map.on("load", (event) => {
            fire(this, "api-load", {});
        });

        this.#map = map;


        return;

        if (!this.hasAttribute("latitude") || !this.hasAttribute("longitude")) {
            return;
        }

        if (!this.#map) {
            
        }



        


       
        //this.onLoad(map);;

    }










    set mark(info) {
        console.log(info)
        let mark = $(this.getMark(info.name));
        if (!mark) {

            if (!info.visible) {
                return;
            }


            mark = $(this).create("mapbox-mark");


        }

        if (!info.visible) {

            mark.remove();
            return;
        }


        mark.prop("update", info);

    }

    set updateMark(info){
        const mark = $(this.getMark(info.name));
        if(!mark){
            return;
        }
        mark.prop("update", info);
    }

    getMark(name) {

        return this.querySelector(`mapbox-mark[name="${name}"]`);
    }

    getApi() {
        return this.#map;
    }

    public panTo(position) {
        const latLng = { lat: position.latitude, lng: position.longitude };
        this.getApi().panTo(latLng);

    }

    public flyTo(info) {


        this.getApi().flyTo({
            center: [info.longitude, info.latitude],
            zoom: info.zoom,
            speed: 3.0,
            curve: 1,
            easing(t) {
                return t;
            }
        });

        console.log(this.getBounds());

    }

    

    getCenter(){
        const center = this.getApi().getCenter();

        return {
            latitude: center.lat,
            longitude: center.lng
        };
    }

    _center(){
        if (!this.hasAttribute("latitude") || !this.hasAttribute("longitude")) {
            this.getApi().center({ lat: Number(this.latitude), lng: Number(this.longitude) });
        }
    }

    getBounds() {
        const bounds = this.getApi().getBounds();
        console.log(bounds)
        return {
            north: bounds.getNorth(),
            south: bounds.getSouth(),
            east: bounds.getEast(),
            west: bounds.getWest(),
        }

    }

    public fitBounds(bounds, padding?) {

        console.log(bounds);

        const coordinates = [
            [bounds.west, bounds.south], // southwestern corner of the bounds
            [bounds.east, bounds.north] // northeastern corner of the bounds
        ];


        console.log(coordinates);
        const zoom = this.getApi().getZoom();

        if (padding === undefined) {
            padding = 0;
        }

        const _bounds = coordinates.reduce(function (bounds, coord) {
            return bounds.extend(coord);
        }, new mapboxgl.LngLatBounds(coordinates[0], coordinates[0]));

        this.getApi().fitBounds(_bounds, {
            padding: padding,
            maxZoom: zoom,
            linear: false
        });
    }


    getZoom() {
        return this.getApi().getZoom();
    }

    setZoom(zoom: number) {
        this.getApi().setZoom(zoom);
    }

}

customElements.define("mapbox-maps", MapboxMaps);
