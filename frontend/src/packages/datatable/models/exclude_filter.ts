// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

export default class ExcludeFilter {
  values: Array<any> = new Array();

  addValue(value: any) {
    this.values.push(value);
  }

  removeValue(value: any) {
    const index = this.values.indexOf(value);
    if (index > -1) {
      this.values.splice(index, 1);
    }
  }
}
