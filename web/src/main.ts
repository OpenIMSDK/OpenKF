/**
 * Copyright © 2023 OpenIM open source community. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';
import router from './router';
import TDesign from 'tdesign-vue-next';
import { store } from './store';
import 'tdesign-vue-next/es/style/index.css';
import './style/index.less';

import { getSDK } from '@/utils/open-im-sdk-wasm';

const OpenIM = getSDK();
console.log(OpenIM);

const app = createApp(App);
app.config.globalProperties.$https = axios; // use axios
app.use(router).use(TDesign).use(store).mount('#app'); // mount the router on the app
