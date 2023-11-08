export class FilterStore {
    private static _store: {[key: string]: (e: any) => boolean } = {};

    public static register<T>(fn: (e: T) => boolean) {
        if (!fn) {
            return;
        }
        const name = fn.name;
        if (!name) {
            return;
        }
        FilterStore._store[name] = fn;
    }

    public static get<T>(name: string): ((e: T) => boolean) | null {
        return FilterStore._store[name];
    }
}