package anchor

import (
	"encoding/binary"
	"math/big"

	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets/anchor/generated/zeta"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// https://github.com/zetamarkets/sdk/blob/79b4f37dea9c494091b924d241c3437885bd6f5f/src/decimal.ts#L23

const (
	SCALE_MASK  = 0x00ff_0000
	SCALE_SHIFT = 16
	SIGN_MASK   = 0x8000_0000
)

func FromAnchorToDecimals(a zeta.AnchorDecimal) float64 {
	// spew.Dump(a)
	scale := scale(a)
	// spew.Dump(scale)
	array := make([]byte, 0, 4*3*4)
	// AppendUint32 does the work for us.
	array = binary.BigEndian.AppendUint32(array, a.Hi)
	array = binary.BigEndian.AppendUint32(array, a.Mid)
	array = binary.BigEndian.AppendUint32(array, a.Lo)

	z := new(big.Int)
	z.SetBytes(array)
	// spew.Dump(array)
	// spew.Dump(z)
	return rainbow.ToFloat(z, int64(scale)-2)
}

func scale(a zeta.AnchorDecimal) uint32 {
	return (a.Flags & SCALE_MASK) >> SCALE_SHIFT
}

// TODO port all the other functions
/****

// Typescript port for rust decimal deserialization.

import { AnchorDecimal } from "./program-types";
import { BN } from "@project-serum/anchor";

const SCALE_MASK: number = 0x00ff_0000;
const SCALE_SHIFT: number = 16;
const SIGN_MASK: number = 0x8000_0000;

export class Decimal {
  private _flags: number;
  private _hi: number;
  private _lo: number;
  private _mid: number;

  public constructor(flags: number, hi: number, lo: number, mid: number) {
    this._flags = flags;
    this._hi = hi;
    this._lo = lo;
    this._mid = mid;
  }

  public static fromAnchorDecimal(decimal: AnchorDecimal): Decimal {
    return new Decimal(decimal.flags, decimal.hi, decimal.lo, decimal.mid);
  }

  public scale(): number {
    return (this._flags & SCALE_MASK) >> SCALE_SHIFT;
  }

  public isSignNegative(): boolean {
    return (this._flags & SIGN_MASK) > 0;
  }

  public isSignPositive(): boolean {
    return (this._flags & SIGN_MASK) == 0;
  }

  public toBN(): BN {
    let bytes = [
      (this._hi >> 24) & 0xff,
      (this._hi >> 16) & 0xff,
      (this._hi >> 8) & 0xff,
      this._hi & 0xff,
      (this._mid >> 24) & 0xff,
      (this._mid >> 16) & 0xff,
      (this._mid >> 8) & 0xff,
      this._mid & 0xff,
      (this._lo >> 24) & 0xff,
      (this._lo >> 16) & 0xff,
      (this._lo >> 8) & 0xff,
      this._lo & 0xff,
    ];

    return new BN(new Uint8Array(bytes));
  }

  public isUnset(): boolean {
    return this._hi == 0 && this._mid == 0 && this._lo == 0 && this._flags == 0;
  }

  public toNumber(): number {
    if (this.isUnset()) {
      return 0;
    }

    let scale = this.scale();
    if (scale == 0) {
      // TODO don't need yet as we don't expect scale 0 decimals.
      throw Error("Scale 0 is not handled.");
    }

    let bn = this.toBN();
    // We use toString because only 53 bits can be stored for floats.
    return (bn.toString() as any) / 10 ** scale;
  }
}
***/
