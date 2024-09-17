export const ICONS = {
  clock: "fa-regular fa-clock",
  play_circle: "fa-regular fa-play-circle",
  send: "fa-regular fa-paper-plane",
  check: "fa-regular fa-circle-check",
  check_solid: "fa-solid fa-circle-check",
  rotate_right: "fa-solid fa-rotate-right",
  arrow_right: "fa-solid fa-arrow-right",
  calender_solid: "fa-solid fa-calendar-days",
  xmark: "fa-solid fa-xmark",
  caret_up: "fa-solid fa-caret-up",
  caret_down: "fa-solid fa-caret-down",
  exclaimation: "fa-solid fa-exclamation-triangle",
  dumbbell: "fa-solid fa-dumbbell",
  arrow_back: "fa-solid fa-arrow-left",
  target: "fa-solid fa-bullseye",
  exclaimation_circle: "fa-solid fa-exclamation-circle",
  xmark_circle: "fa-solid fa-xmark-circle",
  file: "fa-solid fa-file",
  rupee: "fa-solid fa-indian-rupee-sign",
  chevron_down: "fa-solid fa-chevron-down",
  chevron_up: "fa-solid fa-chevron-up",
  chevron_right: "fa-solid fa-chevron-right",
  chevron_left: "fa-solid fa-chevron-left",
  tags: "fa fa-tags",
  undo: "fa-solid fa-undo", // added for REFUNDED
  undo_alt: "fa-solid fa-undo-alt", // added for PARTIALLY_REFUNDED
  hourglass_half: "fa-regular fa-hourglass-half", // added for PENDING
  sync: "fa-solid fa-sync-alt", // added for PROCESSING
  ban: "fa-solid fa-ban", // added for CANCELLED
  thumbs_up: "fa-solid fa-thumbs-up",
  bank: "fa-solid fa-bank",
  ellipses_horizontal: "fa-solid fa-ellipsis",
};
type IconType = keyof typeof ICONS;
export type { IconType };
