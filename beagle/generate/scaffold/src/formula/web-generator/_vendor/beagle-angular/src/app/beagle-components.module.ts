import { NgModule } from '@angular/core'
// import all the components you're going to use with beagle

const components = [
  // your components
]

@NgModule({
  declarations: components,
  entryComponents: components,
  exports: components,
  imports: [
    // everything your components depend on
  ],
})
export class BeagleComponentsModule {}
