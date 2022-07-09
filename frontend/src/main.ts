// Copyright 2021-2022 Teal.Finance contributors
// This file is part of Teal.Finance/Rainbow,
// a screener fo DeFi options, under the MIT License.
// SPDX-License-Identifier: MIT

import { createApp } from 'vue'
import App from './App.vue'
import router from './router';
import './assets/index.css'

createApp(App).use(router).mount('#app')
