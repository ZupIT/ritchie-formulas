import { BeagleModule } from '@zup-it/beagle-angular';
// import all the components you wish to use with Beagle.

@BeagleModule({
  baseUrl: '${bff_url}',
  module: {
    path: './beagle-components.module',
    name: 'BeagleComponentsModule',
  },
  components: {
    // Associate every beagle component to your angular component.
  },
})
export class Beagle {}
