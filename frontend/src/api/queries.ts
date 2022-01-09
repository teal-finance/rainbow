const allQuery = `{
  rows {
    date
    expiry
    provider
    asset
    strike
    call {
      bid {
        px
        size
      }
      ask {
        px
        size
      }
    }
    put {
      bid {
        px
        size
      }
      ask {
        px
        size
      }
    }
  }
}`;

export { allQuery }