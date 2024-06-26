// Websocket
type EventHandler = (event?: any) => void;
export class WebSocket {
    // Ereignishandler
    private onopenHandler: EventHandler | null = null;
    private onmessageHandler: EventHandler | null = null;
    private onerrorHandler: EventHandler | null = null;
    private oncloseHandler: EventHandler | null = null;

    // Konstruktor
    constructor(public url: string) {
        console.log(`WebSocket connection to '${url}' will be simulated.`);
    }

    // Methoden zum Setzen der Ereignishandler
    set onopen(handler: EventHandler) {
        this.onopenHandler = handler;
    }

    set onmessage(handler: EventHandler) {
        this.onmessageHandler = handler;
    }

    set onerror(handler: EventHandler) {
        this.onerrorHandler = handler;
    }

    set onclose(handler: EventHandler) {
        this.oncloseHandler = handler;
    }

    // Methode zum Simulieren des Sendens einer Nachricht
    send(data: string) {
        console.log(`Sending message: ${data}`);
        // Simulieren Sie eine Antwort vom Server nach einer kurzen Verzögerung
        setTimeout(() => {
            this.onmessageHandler?.({ data: `Echo: ${data}` });
        }, 500);
    }

    // Methode zum Simulieren des Öffnens der Verbindung
    open() {
        console.log(`Simulating open WebSocket connection to ${this.url}`);
        this.onopenHandler?.();
    }

    // Methode zum Simulieren des Schließens der Verbindung
    close() {
        console.log(`Closing WebSocket connection to ${this.url}`);
        this.oncloseHandler?.({ code: 1000, reason: "Normal closure" });
    }

    // Methode zum Simulieren eines Fehlers
    error() {
        console.log(`Simulating WebSocket error`);
        this.onerrorHandler?.(new Error("Simulated error"));
    }
}