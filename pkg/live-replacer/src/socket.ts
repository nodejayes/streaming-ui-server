export class Socket {
  constructor(private wsLocation: string) {}

  async openWebSocket(): Promise<WsApi> {
    return new Promise((resolve) => {
      const ws = new WebSocket(`${this.wsLocation}`);
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