import { ClickEventData } from "./types";
import { BaseEvent } from "./base.event";

export class ClickEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as PointerEvent;
        if (!event) {
            console.warn("click event is not of type PointerEvent", e);
            return;
        }
        if (!event.target) {
            console.warn("missing click event target", e);
            return;
        }
        ClickEvent.base?.handleEvent(ClickEvent.typ, event.target as HTMLElement, () => ({
            typ: ClickEvent.typ,
            ctrlKey: event.ctrlKey,
        } as ClickEventData));
    }

    private static get typ() {
        return "click";
    }

    public static register(api: WsApi) {
        ClickEvent.base.reattach(api, ClickEvent.typ, ClickEvent.eventHandler);
    }
}