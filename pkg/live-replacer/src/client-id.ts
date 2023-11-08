export class ClientIdStore {
  constructor(private clientIdKey: string, private identityLocation: string) {}

  getClientId(): string | null {
    return localStorage.getItem(this.clientIdKey);
  }
  
  setClientId(id: string) {
    localStorage.setItem(this.clientIdKey, id);
  }
  
  ensureClientId(): Promise<void> {
    return new Promise((resolve, reject) => {
      const id = this.getClientId();
      if (!id) {
        fetch(this.identityLocation)
          .then((response) => {
            return response.text();
          })
          .then((id) => {
            this.setClientId(id);
            resolve();
          })
          .catch((err) => reject(err));
      } else {
        resolve();
      }
    });
  }
}
