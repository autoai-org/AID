import React from 'react';
import classnames from 'classnames';
import Layout from '@theme/Layout';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './styles.module.css';
import Markdown from '../components/Pages/ExternalMarkdown';

function Changelog() {
  const context = useDocusaurusContext();
  const { siteConfig = {} } = context;
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />">
      <header className={classnames('hero', styles.heroBanner)}>
      <h1 className={classnames('hero__title', styles.title)}>AID Changelog</h1>
      </header>
      <main>
      <div className="container">
        <Markdown src={'https://raw.githubusercontent.com/autoai-org/AID/master/CHANGELOG.md'}></Markdown>
      </div>
      </main>
    </Layout>
  );
}

export default Changelog;
