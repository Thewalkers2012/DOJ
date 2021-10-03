const adminRoutes = [
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('../../views/admin/Admin.vue'),
  },
];

export default adminRoutes;
