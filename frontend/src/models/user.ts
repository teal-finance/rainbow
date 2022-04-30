import { User as SwUser } from "@snowind/state";
import { useStorage } from "@vueuse/core";

export default class User extends SwUser {
  currentPreset = useStorage("preset", "All");
}