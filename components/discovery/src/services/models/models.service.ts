// Initializes the `models` service on path `/models`
import { ServiceAddons } from '@feathersjs/feathers';
import { Application } from '../../declarations';
import { Models } from './models.class';
import createModel from '../../models/models.model';
import hooks from './models.hooks';

// Add this service to the service type index
declare module '../../declarations' {
  interface ServiceTypes {
    'models': Models & ServiceAddons<any>;
  }
}

export default function (app: Application) {
  const options = {
    Model: createModel(app),
    paginate: app.get('paginate')
  };

  // Initialize our service with any options it requires
  app.use('/models', new Models(options, app));

  // Get our initialized service so that we can register hooks
  const service = app.service('models');

  service.hooks(hooks);
}
