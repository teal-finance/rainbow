// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

const addr = import.meta.env.VITE_ADDR as string || "http://localhost:8090"
const serverUrl = addr + import.meta.env.BASE_URL

export { serverUrl }
