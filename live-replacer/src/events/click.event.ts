import {ClickEventData, EventData} from "./types";

export class ClickEvent {
    private static blocked = false;
    private static api: WsApi | null = null;
    private static eventHandler = (e: Event) => {
        if (ClickEvent.blocked) {
            return;
        }
        const event = e as PointerEvent;
        if (!event) {
            console.warn("click event is not of type PointerEvent", e);
            return;
        }
        if (!event.target) {
            console.warn("missing click event target", e);
            return;
        }
        const delayName = `lr${ClickEvent.typ}Delay`;
        const target = event.target as HTMLElement;
        if (!target) {
            console.warn(`no element on target event ${ClickEvent.typ}`);
            return;
        }
        const delay = target.getAttribute(delayName);
        if (delay) {
            const args = delay.split(":");
            let delayTime = 0;
            let delayRun = 0;
            if (args.length > 0) {
                delayTime = isNaN(parseInt(args[0])) ? 0 : parseInt(args[0]);
                if (args.length > 1) {
                    delayRun = isNaN(parseInt(args[1])) ? 0 : parseInt(args[1]);
                }
            }
            ClickEvent.blocked = true;
            setTimeout(() => {
                ClickEvent.blocked = false;
            }, delayTime);
            setTimeout(() => {
                ClickEvent.handle(target, ClickEvent.typ, {
                    typ: ClickEvent.typ,
                    ctrlKey: event.ctrlKey,
                } as ClickEventData);
            }, delayRun);
            return;
        }
        ClickEvent.handle(target, ClickEvent.typ, {
            typ: ClickEvent.typ,
            ctrlKey: event.ctrlKey,
        } as ClickEventData);
    }

    private static handle(target: HTMLElement, eventName: string, eventData: EventData) {
        const actionName = `lr${eventName}action`;
        const payloadName = `lr${eventName}payload`;
        const inputsName = `lr${eventName}inputs`;
        const action = target.getAttribute(actionName);
        if (!action) {
            console.warn(`missing ${actionName} on element`, target);
            return;
        }
        const elementId = target.getAttribute("id") ?? null;
        const payload = target.getAttribute(payloadName) ?? null;
        const inputSelectors: string | null = target.getAttribute(inputsName) ?? null;
        if (!payload && !inputSelectors) {
            ClickEvent.api?.send({
                elementId,
                action: {
                    type: action,
                    payload: null,
                },
                inputs: {},
                eventData,
            });
            return;
        }
        if (!inputSelectors) {
            ClickEvent.api?.send({
                elementId,
                action: {
                    type: action,
                    payload: JSON.parse(payload ?? "null"),
                },
                inputs: {},
                eventData,
            });
            return;
        }

        const selectors = inputSelectors.split("<=>");
        const inputData: { [key: string]: { [key: string]: string } } = {};
        for (const selector of selectors) {
            document.querySelectorAll(selector)?.forEach((el) => {
                el.querySelectorAll("input")?.forEach((input: HTMLInputElement) => {
                    const inputName = input.getAttribute("name");
                    if (!inputName) {
                        return;
                    }
                    if (!inputData[selector]) {
                        inputData[selector] = {};
                    }
                    inputData[selector][inputName] = input.value;
                });
            });
        }
        ClickEvent.api?.send({
            elementId,
            action: {
                type: action,
                payload: JSON.parse(payload ?? "null"),
            },
            inputs: {},
            eventData,
        });
    }

    static get typ() {
        return "click";
    }

    constructor(api: WsApi) {
        ClickEvent.api = api;
        document.querySelectorAll(`[lr${ClickEvent.typ}action]`).forEach((el) => {
            el.removeEventListener(ClickEvent.typ, ClickEvent.eventHandler);
            el.addEventListener(ClickEvent.typ, ClickEvent.eventHandler);
        });
    }
}