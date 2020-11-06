import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
//import Room from './components/Rooms';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    //router.registerRoute('/room', Room);
  },
});