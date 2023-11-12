import { BaseEvent } from "./base.event";
import { BlurEventData } from "./types";

export class BlurEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${BlurEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${BlurEvent.typ} event target`, e);
            return;
        }
        BlurEvent.base?.handleEvent(BlurEvent.typ, event.target as HTMLElement, () => ({
            typ: BlurEvent.typ,
        } as BlurEventData), e);
    };
    private static get typ() {
        return "blur"
    }
    public static register(api: WsApi) {
        BlurEvent.base.reattach(api, BlurEvent.typ, BlurEvent.eventHandler);
    }
}