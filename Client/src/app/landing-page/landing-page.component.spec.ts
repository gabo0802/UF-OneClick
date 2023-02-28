import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatTabContent } from '@angular/material/tabs';
import { MaterialDesignModule } from '../material-design/material-design.module';

import { LandingPageComponent } from './landing-page.component';

describe('LandingPageComponent', () => {
  let component: LandingPageComponent;
  let fixture: ComponentFixture<LandingPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        LandingPageComponent,
        MatTabContent
       ],
       imports: [
        MaterialDesignModule,        
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
