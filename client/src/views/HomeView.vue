<template>
  <div style="padding: 32px; background: linear-gradient(135deg, #f5f7fa, #c3cfe2); min-height: 100vh;">
    <div>
      <h1 style="font-weight: bold;">Bem-vindo ao Gestor de Pokémon e Utilizadores</h1>
      <p>Escolha o serviço, tipo de objeto e ação para interagir com o backend.</p>
    </div>

    <v-card class="pa-4" elevation="3" style="margin-top: 24px;">
      <div style="display: flex; gap: 16px;">
        <v-select
          v-model="selectedService"
          label="Serviço"
          :items="['REST', 'SOAP', 'GraphQL', 'GRPC']"
          variant="solo-filled"
          style="flex: 1"
          @update:modelValue="resetSelections"
        ></v-select>

        <v-select
          v-model="selectedType"
          label="Tipo"
          :items="typeOptions"
          variant="solo-filled"
          style="flex: 1"
          :disabled="!selectedService"
          @update:modelValue="resetAction"
        ></v-select>

        <v-select
          v-model="selectedAction"
          label="Ação"
          :items="actionOptions"
          variant="solo-filled"
          style="flex: 1"
          :disabled="!selectedType"
        ></v-select>
      </div>
    </v-card>

    <div style="margin-top: 32px; min-height: 300px; border: 2px dashed #ccc; border-radius: 8px; padding: 16px;">
      <component
        :is="dynamicComponent"
        :key="componentKey"
        @actionCompleted="handleActionCompleted"
      />
    </div>
  </div>
</template>

<script>
// Importa os componentes específicos para REST
import RestPokemonForm from '@/components/rest/RestPokemonForm.vue';
import RestPokemonList from '@/components/rest/RestPokemonList.vue';
import RestUserForm from '@/components/rest/RestUserForm.vue'; // Precisamos criar este
import RestUserSearch from '@/components/rest/RestUserSearch.vue'; // Precisamos criar este

// Outros views/serviços (manter para a estrutura geral)
// import SoapView from './SoapView.vue';
// import GraphQLView from './GraphQLView.vue';
// import GrpcView from './GrpcView.vue';

export default {
  name: "HomeView",
  components: {
    RestPokemonForm,
    RestPokemonList,
    RestUserForm, // Registra o novo componente
    RestUserSearch, // Registra o novo componente
    // SoapView,
    // GraphQLView,
    // GrpcView
  },
  data() {
    return {
      selectedService: null,
      selectedType: null,
      selectedAction: null,
      componentKey: 0, // Para forçar a recriação do componente

      // Opções de tipo e ação baseadas nos serviços e na documentação
      serviceOptions: {
        'REST': {
          'Pokemons': ['Criar', 'Listar', 'Atualizar', 'Eliminar'], // "Atualizar" e "Eliminar" vão precisar de um ID
          'Utilizadores': ['Criar', 'Buscar por ID', 'Buscar por Username'],
        },
        // 'SOAP': { /* ... */ },
        // 'GraphQL': { /* ... */ },
        // 'GRPC': { /* ... */ },
      }
    };
  },
  computed: {
    typeOptions() {
      // Retorna os tipos disponíveis para o serviço selecionado
      return this.selectedService ? Object.keys(this.serviceOptions[this.selectedService]) : [];
    },
    actionOptions() {
      // Retorna as ações disponíveis para o tipo e serviço selecionados
      if (this.selectedService && this.selectedType) {
        return this.serviceOptions[this.selectedService][this.selectedType] || [];
      }
      return [];
    },
    dynamicComponent() {
      // Determina qual componente exibir com base nas seleções
      if (this.selectedService === 'REST') {
        if (this.selectedType === 'Pokemons') {
          switch (this.selectedAction) {
            case 'Criar':
            case 'Atualizar': // Pode ser o mesmo formulário com lógica condicional
            case 'Eliminar': // Pode ser um formulário para ID ou uma ação na lista
              return 'RestPokemonForm'; // Um formulário genérico para criar/editar/deletar
            case 'Listar':
              return 'RestPokemonList';
          }
        } else if (this.selectedType === 'Utilizadores') {
          switch (this.selectedAction) {
            case 'Criar':
              return 'RestUserForm';
            case 'Buscar por ID':
            case 'Buscar por Username':
              return 'RestUserSearch';
          }
        }
      }
      // Outros serviços (SOAP, GraphQL, GRPC) seriam tratados aqui
      // if (this.selectedService === 'SOAP') { /* ... */ }

      // Default message if nothing is selected or no component matches
      return {
        template: '<span style="color: #888;">Selecione um serviço, tipo e ação para ver os dados ou o formulário.</span>'
      };
    }
  },
  methods: {
    resetSelections() {
      // Reseta tipo e ação quando o serviço muda
      this.selectedType = null;
      this.selectedAction = null;
      this.componentKey++; // Recria o componente dinâmico
    },
    resetAction() {
      // Reseta a ação quando o tipo muda
      this.selectedAction = null;
      this.componentKey++; // Recria o componente dinâmico
    },
    handleActionCompleted() {
      // Método para ser chamado pelos componentes filhos quando uma ação é concluída
      // Por exemplo, quando um Pokémon é criado, queremos atualizar a lista.
      // Poderíamos usar um evento específico ou simplesmente forçar uma atualização.
      this.componentKey++; // Força a recriação do componente atual (se for o de lista, ele irá buscar os dados novamente)
      // Ou poderíamos ter uma lógica mais sofisticada para atualizar apenas o que é necessário.
    }
  }
};
</script>

<style scoped>
/* Adicione estilos específicos para HomeView se necessário */
</style>