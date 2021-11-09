// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'
import User from "./models/user";

const breakpoints = useBreakpoints(breakpointsTailwind)

const isMobile = breakpoints.isSmaller('sm');



const user = new User();

export { user, isMobile }