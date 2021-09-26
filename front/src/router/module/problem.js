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
  {
    path: '/problem/:id/description',
    name: 'description',
    component: () => import('../../components/description/Description.vue'),
  },
];

export default problemRoutes;
