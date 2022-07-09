// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import { User as SwUser } from "@snowind/state";
import { useStorage } from "@vueuse/core";

export default class User extends SwUser {
  currentPreset = useStorage("preset", "All");
}
