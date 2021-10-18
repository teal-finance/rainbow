type StringValidator = (v: string) => boolean;

function guidGenerator(): string {
  const S4 = function () {
    return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
  };
  return (
    S4() +
    S4() +
    "-" +
    S4() +
    "-" +
    S4() +
    "-" +
    S4() +
    "-" +
    S4() +
    S4() +
    S4()
  );
}

const Validators = {
  minStringLength: (len: number): StringValidator => {
    return (v: string): boolean => {
      if (v.length >= len) {
        return true;
      }
      return false;
    }
  }
};

export { guidGenerator, Validators }

