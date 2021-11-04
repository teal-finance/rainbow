// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

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



