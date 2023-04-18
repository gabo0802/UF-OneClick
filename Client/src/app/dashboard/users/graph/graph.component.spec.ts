import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GraphComponent } from './graph.component';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';

describe('GraphComponent', () => {
  let component: GraphComponent;
  let fixture: ComponentFixture<GraphComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        GraphComponent
      ],
      imports: [
        MaterialDesignModule,
        HttpClientModule,
        BrowserAnimationsModule,
        ReactiveFormsModule,
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GraphComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('startMonth should be initially null', ()=> {

    expect(component.datesForm.get('startMonth')?.getRawValue()).toEqual(null);
  });

  it('endMonth should be initially null', ()=> {

    expect(component.datesForm.get('endMonth')?.getRawValue()).toEqual(null);
  });

  it('startYear should be initially null', ()=> {

    expect(component.datesForm.get('startYear')?.getRawValue()).toEqual(null);
  });

  it('endYear should be initially null', ()=> {

    expect(component.datesForm.get('endYear')?.getRawValue()).toEqual(null);
  });

  it('initialState should be initially false', () => {

    expect(component.initialState).toEqual(false);
  });

  it('barChartLegend should initially be false', ()=> {

    expect(component.barChartLegend).toEqual(false);
  });

  it('barChartPlugins should initially be empty []', ()=> {

    expect(component.barChartPlugins).toEqual([]);
  });
});
