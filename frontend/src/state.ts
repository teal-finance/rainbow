// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT
import { ref, reactive } from 'vue';
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'
import User from "@/models/user";
import SwDataTableModel from "@/packages/datatable/models/datatable";
import { OptionsTable } from '@/models/options/types';

const breakpoints = useBreakpoints(breakpointsTailwind);

const isMobile = breakpoints.smaller('sm');
const isTablet = breakpoints.between('sm', 'lg');
const isDesktop = breakpoints.greater('lg');

const user = new User();

const datatable = ref(new SwDataTableModel<OptionsTable>());
const isDatatableReady = ref(false);
const filterConf = reactive<Record<string, Record<string, boolean>>>({
  assets: { 'defaultValue': true },
  providers: { 'defaultValue': true },
});

export { user, isMobile, isTablet, isDesktop, datatable, isDatatableReady, filterConf }
