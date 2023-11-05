import { EventData } from "./types";

export class BaseEvent {
    private blocked: {[key: string]: boolean} = {};
    private api: WsApi | null = null;

    public reattach(api: WsApi, typ: string, eventHandler: (e: Event) => void) {
        this.api = api;
        document.querySelectorAll(`[lr${typ}action]`).forEach((el) => {
            el.removeEventListener(typ, eventHandler);
            el.addEventListener(typ, eventHandler);
        });
    }

    public handleEvent(typ: string, target: HTMLElement, eventDataBuilder: () => EventData) {
        if (!target) {
            console.warn(`no element on target event ${typ}`);
            return;
        }
        const elementId = target.getAttribute("id") ?? '';
        if (this.blocked[elementId]) {
            return;
        }

        // handle Delay
        const delayName = `lr${typ}Delay`;
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
            this.blocked[elementId] = true;
            setTimeout(() => {
                this.blocked[elementId] = false;
            }, delayTime);
            setTimeout(() => {
                this.handle(target, typ, eventDataBuilder(), elementId);
            }, delayRun);
            return;
        }

        // run without Delay
        this.handle(target, typ, eventDataBuilder(), elementId);
    }

    private handle(target: HTMLElement, eventName: string, eventData: EventData, elementId: string) {
        const actionName = `lr${eventName}action`;
        const payloadName = `lr${eventName}payload`;
        const inputsName = `lr${eventName}inputs`;
        const action = target.getAttribute(actionName);
        if (!action) {
            console.warn(`missing ${actionName} on element`, target);
            return;
        }
        const payload = target.getAttribute(payloadName) ?? null;
        const inputSelectors: string | null = target.getAttribute(inputsName) ?? null;
        if (!payload && !inputSelectors) {
            this.api?.send({
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
            this.api?.send({
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
        this.api?.send({
            elementId,
            action: {
                type: action,
                payload: JSON.parse(payload ?? "null"),
            },
            inputs: {},
            eventData,
        });
    }
}