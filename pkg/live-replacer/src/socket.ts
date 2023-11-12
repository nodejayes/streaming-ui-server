import { ClientIdStore } from "./client-id";

export class Socket {
  constructor(private wsLocation: string, private clientIdStore: ClientIdStore) {}

  async openWebSocket(): Promise<WsApi> {
    return new Promise((resolve) => {
      const ws = new WebSocket(`${this.wsLocation}?clientId=${this.clientIdStore.getClientId()}&pageId=${localStorage.getItem('pageId')}`);
      const api = {
        send: (data: any) => ws.send(JSON.stringify(data)),
        onMessage: (_ev: MessageEvent<any>) => {},
        onClose: (_ev: CloseEvent) => {},
      };
      ws.onopen = () => {
        resolve(api);
      };
      ws.onmessage = (ev) => api.onMessage(ev);
      ws.onclose = (ev) => api.onClose(ev);
    });
  }
}