import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import RestView from '../views/RestView.vue';
import SoapView from '../views/SoapView.vue';
import GraphQLView from '../views/GraphQLView.vue';
import GrpcView from '../views/GrpcView.vue';

const routes = [
  { 
    path: '/', 
    name: 'Home', 
    component: HomeView 
  },
  { 
    path: '/rest', 
    name: 'Rest', 
    component: RestView 
  },
  { 
    path: '/soap', 
    name: 'Soap', 
    component: SoapView 
  },
  { 
    path: '/graphql', 
    name: 'GraphQL', 
    component: GraphQLView 
  },
  { 
    path: '/grpc', 
    name: 'Grpc', 
    component: GrpcView 
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
