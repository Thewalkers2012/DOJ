const ContextRoutes = [
  {
    path: '/context',
    name: 'ContextList',
    component: () => import('../../components/contextList/ContextList.vue'),
  },
  {
    path: '/context/:id',
    name: 'Context',
    component: () => import('../../components/context/Context.vue'),
  },
];

export default ContextRoutes;
