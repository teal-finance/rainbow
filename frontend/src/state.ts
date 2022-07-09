// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'
import User from "@/models/user";

const breakpoints = useBreakpoints(breakpointsTailwind);

const isMobile = breakpoints.smaller('sm');
const isTablet = breakpoints.between('sm', 'lg');
const isDesktop = breakpoints.greater('lg');

const user = new User();

export { user, isMobile, isTablet, isDesktop }
