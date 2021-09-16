const problemRoutes = [
  {
    path: '/problem',
    name: 'problem',
    component: () => import('../../components/problemList.vue'),
  },
];

export default problemRoutes;
