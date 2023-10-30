interface Action<T> {
  type: string;
  payload?: T;
  inputs?: { [key: string]: { [key: string]: string } };
}

interface WsApi {
  send: (data: any) => void;
  onMessage: ((msg: MessageEvent<any>) => any) | null;
  onClose: ((e: CloseEvent) => any) | null;
}
