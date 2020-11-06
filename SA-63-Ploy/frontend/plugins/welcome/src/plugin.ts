import { createPlugin } from '@backstage/core';
import Books from './components/Books';
import Login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/Books', Books);
  },
});