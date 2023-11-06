import { EventData } from "./types";
import {FilterStore} from "../filter";

interface ApiMessage {
    elementId: string;
    action: {
        type: string;
    },
    inputs?: {[key: string]: {[key: string]: string}},
    eventData?: EventData,
}

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

    public handleEvent(typ: string, target: HTMLElement, eventDataBuilder: () => EventData, originalEvent: any) {
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
                this.handle(target, typ, eventDataBuilder(), elementId, originalEvent);
            }, delayRun);
            return;
        }

        // run without Delay
        this.handle(target, typ, eventDataBuilder(), elementId, originalEvent);
    }

    private handle(target: HTMLElement, eventName: string, eventData: EventData, elementId: string, originalEvent: any) {
        const actionName = `lr${eventName}action`;
        const action = target.getAttribute(actionName);
        if (!action) {
            console.warn(`missing ${actionName} on element`, target);
            return;
        }

        const msg = this.buildMessage(target, eventName, elementId, action, eventData, originalEvent);
        if (!msg) {
            return;
        }
        this.api?.send(msg);
    }

    private buildMessage(target: HTMLElement, eventName: string, elementId: string, action: string, eventData: EventData, originalEvent: any): ApiMessage | null {
        const inputsName = `lr${eventName}inputs`;
        const filterName = `lr${eventName}Filter`;
        const filterActionName = `lr${eventName}FilterAction`;

        const msg: ApiMessage = {
            elementId,
            action: {
                type: action,
            },
            inputs: {},
            eventData,
        };

        const inputSelectors: string | null = target.getAttribute(inputsName) ?? null;
        if (inputSelectors) {
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
            msg.inputs = inputData;
        }

        const filterFn = target.getAttribute(filterName);
        if (filterFn && typeof FilterStore.get(filterFn) === "function") {
            const filterAction = target.getAttribute(filterActionName);
            const fn = FilterStore.get(filterFn);
            if (fn ? fn(originalEvent) : false) {
                if (filterAction) {
                    msg.action.type = filterAction;
                }
            } else {
                if (!filterAction) {
                    return null;
                }
            }
        }

        return msg;
    }
}