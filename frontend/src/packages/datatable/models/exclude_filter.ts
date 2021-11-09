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