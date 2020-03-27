import React from 'react';
import classnames from 'classnames';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';
import styles from './styles.module.css';
import OpenSource from '../components/Landing/OpenSource';

const features = [
  {
    title: <>Easy to Use</>,
    description: (
      <>
        A.I.D was designed from the ground up to be easily installed and
        used to get your service up and running quickly.
      </>
    ),
  },
  {
    title: <>Focus on What Matters</>,
    description: (
      <>
        A.I.D lets you focus on your algorithm/model, and we&apos;ll do the chores.
      </>
    ),
  },
  {
    title: <>Extensible</>,
    description: (
      <>
        Extend or customize your workflow by install third-party packages and plugins, including from image analysis to model interpretation.
      </>
    ),
  },
];

function Feature({imageUrl, title, description}) {
  const imgUrl = useBaseUrl(imageUrl);
  return (
    <div className={classnames('col col--4', styles.feature)}>
      {imgUrl && (
        <div className="text--center">
          <img className={styles.featureImage} src={imgUrl} alt={title} />
        </div>
      )}
      <h3>{title}</h3>
      <p>{description}</p>
    </div>
  );
}

function Home() {
  const context = useDocusaurusContext();
  const {siteConfig = {}} = context;
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />">
      <header className={classnames('hero hero--primary', styles.heroBanner)}>
        <div className="container">
          <h1 className="hero__title">{siteConfig.title}</h1>
          <p className="hero__subtitle">{siteConfig.tagline}</p>
          <div className={styles.buttons}>
            <Link
              className={classnames(
                'button button--outline button--secondary button--lg',
                styles.getStarted,
              )}
              to={useBaseUrl('docs/quickstart')}>
              Get Started
            </Link>
          </div>
        </div>
      </header>
      <main>
        <div className="container">
          <OpenSource></OpenSource>
        </div>
        {features && features.length && (
          <section className={styles.features}>
            <div className="container">
              <div className="row">
                {features.map((props, idx) => (
                  <Feature key={idx} {...props} />
                ))}
              </div>
            </div>
          </section>
        )}
      </main>
    </Layout>
  );
}

export default Home;
