import { useStorage } from '@vueuse/core'

export default class User {
  name = useStorage<string>("name", "anonymous");
  isDarkMode = useStorage<boolean>("isDarkMode", false);

  constructor() {
    this.checkDarkMode()
  }

  toggleDarkMode() {
    console.log("Toggle dark mode", this.isDarkMode.value, "=>", !this.isDarkMode.value)
    this.isDarkMode.value = !this.isDarkMode.value;
    this.checkDarkMode()
  }

  private checkDarkMode() {
    if (this.isDarkMode.value === true) {
      document.body.classList.add("bg-background-dark", "text-foreground-dark");
    } else {
      document.body.classList.remove("bg-background-dark", "text-foreground-dark");
    }
  }
}



