import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn'
import Checkout from './components/Checkout'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/Checkout', Checkout);
  },
});
