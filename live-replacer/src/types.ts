interface Action<T> {
  type: string;
  payload: T;
}

interface WsApi {
  send: (data: any) => void;
  onMessage: ((msg: MessageEvent<any>) => any) | null;
  onClose: ((e: CloseEvent) => any) | null;
}
