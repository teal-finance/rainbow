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
        iv
      }
      ask {
        px
        size
        iv
      }
    }
    put {
      bid {
        px
        size
        iv
      }
      ask {
        px
        size
        iv
      }
    }
  }
}`;

export { allQuery }