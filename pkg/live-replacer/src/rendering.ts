import { ClickEvent } from "./events/click.event";
import { MouseEnterEvent } from "./events/mouseenter.event";
import { MouseLeaveEvent } from "./events/mouseleave.event";
import { MouseMoveEvent } from "./events/mousemove.event";
import { MouseDownEvent } from "./events/mousedown.event";
import { MouseUpEvent } from "./events/mouseup.event";
import { MouseOutEvent } from "./events/mouseout.event";
import { KeyDownEvent } from "./events/keydown.event";
import { KeyUpEvent } from "./events/keyup.event";

export async function render(api: WsApi) {
    ClickEvent.register(api);
    MouseEnterEvent.register(api);
    MouseLeaveEvent.register(api);
    MouseMoveEvent.register(api);
    MouseDownEvent.register(api);
    MouseUpEvent.register(api);
    MouseOutEvent.register(api);
    KeyDownEvent.register(api);
    KeyUpEvent.register(api);
}