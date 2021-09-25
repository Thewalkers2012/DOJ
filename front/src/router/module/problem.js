const problemRoutes = [
  {
    path: '/problems',
    name: 'problemList',
    component: () => import('../../components/problemList.vue'),
  },
  {
    path: '/problem/:id',
    name: 'problem',
    component: () => import('../../components/problem.vue'),
  },
];

export default problemRoutes;
